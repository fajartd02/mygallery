import Link from "next/link";
import styles from "./Styles/style.module.css";

export default function NavButton({ url, background, color, children }) {
  return (
    <Link
      href={url}
      style={{
        background: background ? background : "inherit",
        color: color ? color : "inherit",
      }}
      className={styles.nav_button}
    >
      {children}
    </Link>
  );
}
