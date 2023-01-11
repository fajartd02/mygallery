import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import Image from "next/image";
import { memo, useContext, useRef, useState } from "react";
import styles from "../Auth/Styles/LoginForm.module.css";
import { toast } from "react-toastify";
import Router from "next/router";
import Spinner from "react-bootstrap/Spinner";
import { MemoryContext } from "../../context/MemoryContextProvider";
import { postFormDataWithCreds } from "../../helper/options";

export default function EditMemoryForm({ show, setShow }) {
  const { memory, setMemory } = useContext(MemoryContext);
  const oldImage = `${process.env.NEXT_PUBLIC_STATIC_IMAGE_URL}/${memory.imageUrl}`;
  const [imageUrl, setImageUrl] = useState("");
  const [fileInput, setFileInput] = useState(null);
  const [progress, setProgress] = useState(false);
  const formRef = useRef(null);

  function resetImage() {
    setImageUrl("");
    setFileInput(null);
  }

  function resetForm() {
    resetImage();
    formRef.current.reset();
  }

  function handleClose() {
    resetForm();
    setShow(false);
  }

  function progressDone() {
    setProgress(false);
    handleClose();
  }

  function changeImage(event) {
    setFileInput(event.target.files[0]);
    if (event.target.files && event.target.files[0]) {
      setImageUrl(URL.createObjectURL(event.target.files[0]));
    }
  }

  async function editMemory(event) {
    event.preventDefault();
    setProgress(true);

    const target = event.target;
    const newTag = target.tag.value;
    const newDescription = target.description.value;

    const formData = new FormData();

    let changed = false;

    if (newTag != memory.tag) {
      formData.append("tag", target.tag.value);
      changed = true;
    }

    if (newDescription != memory.description) {
      formData.append("description", target.description.value);
      changed = true;
    }

    if (fileInput) {
      formData.append("file", fileInput);
      changed = true;
    }

    if (!changed) {
      progressDone();
      return;
    }

    try {
      const res = await fetch(
        `${process.env.NEXT_PUBLIC_URL}/memories/${memory.ID}`,
        postFormDataWithCreds(formData, "PUT")
      );

      const success = res.status == 200;

      if (success) {
        Router.reload(window.location.pathname);
      } else {
        toast.error(`Unable to update... 🤯!`);
      }
    } catch (err) {
      toast.error(`Something went wrong... 🤯!`);
    } finally {
      progressDone();
    }
  }

  return (
    <>
      <Modal
        show={show}
        onHide={handleClose}
        backdrop="static"
        keyboard={false}
      >
        <Modal.Header closeButton>
          <Modal.Title>Edit memory</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form id="add-memo-form" onSubmit={editMemory} ref={formRef}>
            {!fileInput && (
              <Form.Group
                controlId="exampleForm.ControlInput1"
                className="mb-3"
              >
                <div
                  style={{
                    position: "relative",
                    width: "120px",
                    height: "100px",
                  }}
                >
                  <Image
                    unoptimized={() => oldImage}
                    src={oldImage}
                    alt="Unsuported File"
                    fill
                    style={{ objectFit: "contain", background: "#d4d4d4" }}
                  />
                </div>
              </Form.Group>
            )}
            {fileInput && (
              <Form.Group
                controlId="exampleForm.ControlInput1"
                className="mb-3"
              >
                <div
                  style={{
                    position: "relative",
                    width: "120px",
                    height: "100px",
                  }}
                >
                  <Image
                    loader={() => imageUrl}
                    src={imageUrl}
                    alt="Unsuported File"
                    fill
                    style={{ objectFit: "contain", background: "#d4d4d4" }}
                  />
                </div>
              </Form.Group>
            )}
            <Form.Group controlId="formFile" className="mb-3">
              <Form.Label>Image</Form.Label>
              <Form.Control
                type="file"
                onChange={changeImage}
                name="file"
                accept="image/png, image/gif, image/jpeg"
              />
            </Form.Group>
            <Form.Group className="mb-3" controlId="exampleForm.ControlInput1">
              <Form.Label>Tag</Form.Label>
              <Form.Control
                type="tag"
                placeholder="Tag..."
                name="tag"
                defaultValue={memory.tag}
                required
              />
            </Form.Group>
            <Form.Group
              className="mb-3"
              controlId="exampleForm.ControlTextarea1"
            >
              <Form.Label>Description</Form.Label>
              <Form.Control
                as="textarea"
                rows={3}
                name="description"
                placeholder="Description..."
                defaultValue={memory.description}
                required
              />
            </Form.Group>
          </Form>
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={resetForm}>
            Reset
          </Button>
          <Button variant="primary" type="submit" form="add-memo-form">
            Add
          </Button>
        </Modal.Footer>
        {progress && (
          <div
            className={`${styles.overlay} d-flex justify-content-center align-items-center`}
          >
            <Spinner animation="border" role="status">
              <span className="visually-hidden">Loading...</span>
            </Spinner>
          </div>
        )}
      </Modal>
    </>
  );
}
