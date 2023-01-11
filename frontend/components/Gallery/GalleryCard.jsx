import styles from "./Styles/GalleryCard.module.css";
import MemoryImage from "./MemoryImage";
import Router from "next/router";

export default function GalleryCard({ memoryID, date, tag, imageUrl }) {
  return (
    <div className="col d-flex justify-content-center align-items-center">
      <div className={styles.card_container}>
        <div className={styles.card_image}>
          <MemoryImage
            fill
            style={{
              objectFit: "cover",
            }}
            onClick={() => Router.push(`/gallery/${memoryID}`)}
            imageUrl={imageUrl}
          />
        </div>
        <div className={styles.card_body}>
          {date}, {tag}
        </div>
      </div>
    </div>
  );
}
