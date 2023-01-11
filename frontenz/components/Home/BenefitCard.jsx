import styles from "./Styles/BenefitSection.module.css";

export default function BenefitCard({ title, description, children }) {
  return (
    <div className={`col p-md-4 p-2 ${styles.card_container}`}>
      <div className={`d-flex flex-column p-4 ${styles.card}`}>
        <div className={styles.card_icon}>{children}</div>
        <div className={styles.card_title}>{title}</div>
        <div className={styles.card_description}>{description}</div>
      </div>
    </div>
  );
}
