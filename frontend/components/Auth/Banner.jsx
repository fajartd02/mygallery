import styles from "./Styles/Banner.module.css";
import Image from "next/image";

export default function Banner() {
  return (
    <div className={styles.banner_section}>
      <div className={styles.banner_container}>
        <Image
          src="/Assets/securelogin.png"
          alt="banner"
          fill
          sizes="100%"
          priority
        ></Image>
      </div>
    </div>
  );
}
