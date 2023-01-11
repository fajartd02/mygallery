import { useState, useEffect } from "react";

import styles from "./Styles/style.module.css";
import headerStyles from "../Layouts/Styles/Header.module.css";

export default function Hamburger({ navRef }) {
  const [isActive, setActive] = useState(false);

  function HamburgerHandler() {
    setActive(!isActive);
    navRef.current.classList.toggle(headerStyles.collapse);
  }

  return (
    <div
      onClick={HamburgerHandler}
      className={`${styles.hamburger} ${isActive ? styles.active : ""}`}
    >
      <span className={styles.bar}></span>
      <span className={styles.bar}></span>
      <span className={styles.bar}></span>
    </div>
  );
}
