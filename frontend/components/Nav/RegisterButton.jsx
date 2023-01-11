import Link from "next/link";
import styles from "./Styles/style.module.css";

export default function RegisterButton() {
  return (
    <Link
      href="/register"
      className={`${styles.register_btn} ${styles.nav_button}`}
    >
      Register
    </Link>
  );
}
