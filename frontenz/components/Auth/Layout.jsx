import styles from "./Styles/Layout.module.css";

export default function Layout({ children }) {
  return <div className={styles.auth_layout}>{children}</div>;
}
