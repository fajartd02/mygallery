import styles from "./Styles/style.module.css";
import Link from "next/link";

export default function LoginButton() {
  return (
    <Link href="/login" className={`${styles.login_btn} ${styles.nav_button}`}>
      Login
    </Link>
  );
}
