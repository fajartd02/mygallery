import Header from "../../components/Layouts/LoggedHeader";
import MemoryDetail from "../../components/Gallery/MemoryDetail";
import { useEffect, useState } from "react";
import { redirectIfMissingCreds } from "../../helper/auth";
import { MemoryContextProvider } from "../../context/MemoryContextProvider";

export async function getServerSideProps({ params }) {
  return {
    props: {
      id: params.id,
    },
  };
}

export default function GalleryDetails({ id }) {
  useEffect(() => {
    redirectIfMissingCreds();
  }, []);

  return (
    <>
      <MemoryContextProvider>
        <Header />
        <MemoryDetail id={id} />
      </MemoryContextProvider>
    </>
  );
}
