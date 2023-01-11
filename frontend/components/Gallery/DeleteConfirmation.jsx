import { useContext } from "react";
import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import { toast } from "react-toastify";
import { MemoryContext } from "../../context/MemoryContextProvider";
import { requestWithCreds } from "../../helper/options";
import Router from "next/router";

export default function DeleteConfirmation({ show, setShow }) {
  const handleClose = () => setShow(false);
  const { memory, setMemory } = useContext(MemoryContext);

  async function confirmHandler() {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_URL}/memories/${memory.ID}`,
      requestWithCreds("DELETE")
    );

    const resJson = await res.json();
    const success = res.status;

    if (success) {
      toast.success(`${resJson.message} 👌!`);
      setMemory(null);
      Router.replace("/gallery");
    } else {
      toast.error(`Unable to delete memory 🤯!`);
    }
  }

  return (
    <Modal show={show} onHide={handleClose} backdrop="static" keyboard={false}>
      <Modal.Header closeButton>
        <Modal.Title>Delete Confirmation</Modal.Title>
      </Modal.Header>
      <Modal.Body>Are you sure you want to delete this memory ?</Modal.Body>
      <Modal.Footer>
        <Button variant="danger" onClick={handleClose}>
          No
        </Button>
        <Button variant="success" onClick={confirmHandler}>
          Yes
        </Button>
      </Modal.Footer>
    </Modal>
  );
}
