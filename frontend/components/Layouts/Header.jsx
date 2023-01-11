import styles from "./Styles/Header.module.css";
import { useRef } from "react";

import LoginButton from "../Nav/LoginButton";
import RegisterButton from "../Nav/RegisterButton";
import HomeButton from "../Nav/HomeButton";
import Hamburger from "../Nav/Hamburger";
import Link from "next/link";

export default function Header() {
  const navRef = useRef();

  return (
    <>
      <header className={styles.header}>
        <div className={styles.wrapper}>
          <div className={styles.logo}>MyGallery.</div>
          <div className={styles.nav_container}>
            <HomeButton />
            <LoginButton />
            <RegisterButton />
          </div>
          <Hamburger navRef={navRef} />
        </div>
      </header>
      <div className={styles.nav_collapse} ref={navRef}>
        <Link href="/">Beranda</Link>
        <Link href="/login">Login</Link>
        <Link href="/register">Register</Link>
      </div>
    </>
  );
}
