import styles from "./Styles/SignInSection.module.css";
import Link from "next/link";

export default function SignInSection() {
  return (
    <section className={styles.section}>
      <div className={`container ${styles.container} p-5`}>
        <div
          className="d-flex flex-column justify-content-center align-items-center p-5"
          style={{ gap: "30px" }}
        >
          <div className={styles.title}>
            Are you ready to use <br /> this application?
          </div>
          <div className={styles.description}>
            Amet minim mollit non deserunt ullamco est sit aliqua dolor do
          </div>
          <Link className={styles.button} href="/login">
            Let&apos;s Sign In!
          </Link>
        </div>
      </div>
    </section>
  );
}
