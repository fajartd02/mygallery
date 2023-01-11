import styles from "./Styles/MainLayout.module.css";

export default function MainLayout({ children, background }) {
  return (
    <main
      style={{ background: background ? background : "#f5f5f5" }}
      className={styles.main}
    >
      {children}
    </main>
  );
}
