import styles from "./Styles/MemoryUpdateButton.module.css";

export default function MemoryUpdateButton({ action, background, handler }) {
  return (
    <button
      className={styles.container}
      onClick={handler}
      style={{ background: background ? background : "#82AAE3" }}
    >
      <span>{action}</span>
    </button>
  );
}
