import styles from "./Styles/MemoryDetail.module.css";

import Link from "next/link";
import MemoryImage from "./MemoryImage";
import TagBadge from "./TagBadge";
import UpdateButton from "./MemoryUpdateButton";
import { useContext, useEffect, useState } from "react";
import Loading from "../Layouts/Loading";
import { requestWithCreds } from "../../helper/options";
import { MemoryContext } from "../../context/MemoryContextProvider";
import { toast } from "react-toastify";
import Router from "next/router";
import { getFormattedDate, getFormattedTime } from "../../helper/utils";
import EditMemoryForm from "./EditMemoryFormModal";
import DeleteConfirmation from "./DeleteConfirmation";

export default function MemoryDetail({ id }) {
  const [loading, setLoading] = useState(true);
  const [showEditModal, setShowEditModal] = useState(false);
  const [showDeletModal, setShowDeleteModal] = useState(false);
  const { memory, setMemory } = useContext(MemoryContext);

  const handleShowEditModal = () => setShowEditModal(true);
  const handleShowDeleteModal = () => setShowDeleteModal(true);

  useEffect(() => {
    async function fetchMemory() {
      try {
        const res = await fetch(
          `${process.env.NEXT_PUBLIC_URL}/memories/${id}`,
          requestWithCreds()
        );
        const resJson = await res.json();
        const success = res.status == 200;
        if (success) {
          setMemory(resJson.data);
        } else {
          toast.error("Failed to fetch memory 💀");
          Router.replace("/gallery");
        }
      } catch (err) {
        toast.error("Something went wrong 😕");
        Router.replace("/gallery");
      } finally {
        setLoading(false);
      }
    }
    fetchMemory();
  }, [id, setMemory]);

  if (loading || !memory) {
    return <Loading></Loading>;
  }

  return (
    <div className={`container my-4 ${styles.container}`}>
      <div className={`row mb-2 px-3 ${styles.detail_header}`}>
        <Link className={styles.back_button} href="/gallery">
          Back
        </Link>
        <div className={styles.memory_title}>
          {getFormattedDate(memory.CreatedAt)},{" "}
          {getFormattedTime(memory.CreatedAt)}
        </div>
      </div>
      <div
        className="row flex-column justify-content-center align-items-center mb-3 mb-sm-4 px-4"
        style={{ gap: "20px" }}
      >
        <div className={styles.memory_image_container}>
          <MemoryImage
            fill
            imageUrl={memory.imageUrl}
            style={{ objectFit: "contain" }}
          />
        </div>
        <TagBadge tag={memory.tag}></TagBadge>
        <div className="row">
          <div
            className="col-xl-8 col-lg-10 col-12 px-4 m-auto fs-5 text-center"
            style={{ textAlign: "justify" }}
          >
            {memory.description}
          </div>
        </div>
      </div>
      <div className={`row justify-content-center ${styles.detail_footer}`}>
        <UpdateButton
          action="Edit"
          handler={handleShowEditModal}
        ></UpdateButton>
        <EditMemoryForm show={showEditModal} setShow={setShowEditModal} />
        <UpdateButton
          action="Delete"
          background="#E39982"
          handler={handleShowDeleteModal}
        ></UpdateButton>
        <DeleteConfirmation
          show={showDeletModal}
          setShow={setShowDeleteModal}
        />
      </div>
    </div>
  );
}
