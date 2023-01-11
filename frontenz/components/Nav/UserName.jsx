import styles from "./Styles/style.module.css";

export default function UserName({ name }) {
  return (
    <div
      style={{ color: "black", cursor: "auto" }}
      className={styles.nav_button}
    >
      {name}
    </div>
  );
}
