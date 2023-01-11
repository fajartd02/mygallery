import Button from "react-bootstrap/Button";
import Modal from "react-bootstrap/Modal";
import Form from "react-bootstrap/Form";
import { Col, Row } from "react-bootstrap";
import { useContext, useRef, useState } from "react";
import { postJsonWithCreds } from "../../helper/options";
import { toast } from "react-toastify";
import styles from "../Auth/Styles/LoginForm.module.css";
import Spinner from "react-bootstrap/Spinner";
import { MemoryContext } from "../../context/MemoryContextProvider";

export default function SortFilterModal({ show, setShow }) {
  const [isSortNone, setSortNone] = useState(true);
  const [isFilterNone, setFilterNone] = useState(true);
  const [progress, setProgress] = useState(false);
  const formRef = useRef(null);

  const { setSorted, setMemories } = useContext(MemoryContext);

  function resetState() {
    setSortNone(true);
    setFilterNone(true);
  }

  function resetForm() {
    resetState();
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

  function sortChange(e) {
    if (e.target.value === "none") {
      setSortNone(true);
    } else {
      setSortNone(false);
    }
  }

  function filterChange(e) {
    if (e.target.value === "none") {
      setFilterNone(true);
    } else {
      setFilterNone(false);
    }
  }

  async function ApplySortFilter(e) {
    e.preventDefault();
    setProgress(true);

    const target = e.target;
    const payload = {};

    const sortBy = target.sortBy.value;
    const filterBy = target.filterBy.value;

    if (isFilterNone && isSortNone) {
      setSorted(false);
      progressDone();
      return;
    }

    if (!isSortNone) {
      payload["sort"] = {
        by: sortBy,
        order: target.order.value,
      };
    }

    if (!isFilterNone) {
      payload["filter"] = {
        by: filterBy,
        keyword: target.keyword.value,
      };
    }

    try {
      const res = await fetch(
        `${process.env.NEXT_PUBLIC_URL}/memories/list`,
        postJsonWithCreds(payload)
      );

      const resJson = await res.json();
      const success = res.status == 200;

      if (success) {
        setMemories(resJson.data);
        setSorted(true);
        toast.success("Sort & Filter succesfully applied 👌");
      } else {
        toast.error(`Something went wrong 🤯!`);
      }
    } catch (err) {
      toast.error(`Something went wrong 🤯!`);
    } finally {
      progressDone();
    }
  }

  return (
    <Modal show={show} onHide={handleClose} backdrop="static" keyboard={false}>
      <Modal.Header closeButton>
        <Modal.Title>Sort & Filter</Modal.Title>
      </Modal.Header>
      <Modal.Body>
        <Form onSubmit={ApplySortFilter} id="sort-filter-form" ref={formRef}>
          <Form.Group className="mb-5">
            <Form.Label className="fs-5 fw-bold mb-3">Sort</Form.Label>
            <Form.Group as={Row} className="mb-2">
              <Form.Label column sm={2}>
                By
              </Form.Label>
              <Col sm={10}>
                <Form.Select
                  defaultValue="none"
                  name="sortBy"
                  onChange={sortChange}
                >
                  <option value="none">None</option>
                  <option value="created_at">Uploaded time</option>
                  <option value="tag">Tag</option>
                  <option value="updated_at">Last Edited</option>
                </Form.Select>
              </Col>
            </Form.Group>
            <Form.Group as={Row} className="mb-2">
              <Form.Label column sm={2}>
                Order
              </Form.Label>
              <Col sm={10}>
                <Form.Select name="order" disabled={isSortNone}>
                  <option value="asc">Ascending</option>
                  <option value="desc">Descending</option>
                </Form.Select>
              </Col>
            </Form.Group>
          </Form.Group>
          <Form.Group className="mb-3">
            <Form.Label className="fs-5 fw-bold mb-3">Filter</Form.Label>
            <Form.Group as={Row} className="mb-2">
              <Form.Label column sm={2}>
                By
              </Form.Label>
              <Col sm={10}>
                <Form.Select
                  defaultValue="none"
                  name="filterBy"
                  onChange={filterChange}
                >
                  <option value="none">None</option>
                  <option value="tag">Tag</option>
                  <option value="description">Description</option>
                </Form.Select>
              </Col>
            </Form.Group>
            <Form.Group as={Row} className="mb-2">
              <Form.Label column sm={2}>
                Keyword
              </Form.Label>
              <Col sm={10}>
                <Form.Control
                  type="text"
                  placeholder=""
                  name="keyword"
                  disabled={isFilterNone}
                />
              </Col>
            </Form.Group>
          </Form.Group>
        </Form>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={resetForm}>
          Reset
        </Button>
        <Button variant="primary" form="sort-filter-form" type="submit">
          Apply
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
  );
}
