import Header from "../../components/Layouts/LoggedHeader";
import GalleryGridContainer from "../../components/Gallery/GalleryGridContainer";
import ButtonSection from "../../components/Gallery/ButtonSection";
import { useEffect } from "react";
import { redirectIfMissingCreds } from "../../helper/auth";
import { MemoryContextProvider } from "../../context/MemoryContextProvider";

export default function Gallery() {
  useEffect(() => {
    redirectIfMissingCreds();
  }, []);

  return (
    <>
      <Header />
      <MemoryContextProvider>
        <ButtonSection />
        <GalleryGridContainer />
      </MemoryContextProvider>
    </>
  );
}
