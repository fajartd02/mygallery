import Link from "next/link";
import { useState } from "react";
import { postJsonOpt } from "../../helper/options";
import { toast } from "react-toastify";
import Router from "next/router";
import styles from "./Styles/LoginForm.module.css";
import { setUserData } from "../../helper/auth";

export default function LoginForm() {
  const [progress, setProgress] = useState(false);

  async function submitHandler(e) {
    e.preventDefault();
    setProgress(true);

    const data = {
      email: e.target.email.value,
      password: e.target.password.value,
    };

    try {
      let res = await fetch(
        `${process.env.NEXT_PUBLIC_URL}/login`,
        postJsonOpt(data)
      );

      const resData = await res.json();
      const success = resData?.data && true;

      if (success) {
        toast.success(`Welcome back 👋!`);
        const userData = resData.data;
        const userCreds = {
          email: userData.email,
          userID: userData.userID,
          fullname: userData.fullName,
        };
        setUserData(userData.token, userCreds);
        Router.replace("/gallery");
      } else {
        toast.error(`${resData.message} 🤯!`);
      }
    } catch (err) {
      toast.error("Something went wrong...");
    } finally {
      setProgress(false);
      e.target.reset();
    }
  }

  return (
    <form
      onSubmit={submitHandler}
      className="d-flex justify-content-center align-items-center"
    >
      <div className={styles.form_container}>
        {progress && <div className={styles.overlay}></div>}
        <h1 className={styles.form_title}>Sign In</h1>
        <div className={styles.form_input_container}>
          <label htmlFor="email">Email</label>
          <input
            id="email"
            type="email"
            placeholder="Email..."
            name="email"
            required
          />
        </div>
        <div className={styles.form_input_container}>
          <label htmlFor="password">Password</label>
          <input
            id="password"
            type="password"
            placeholder="Password..."
            name="password"
            required
          />
        </div>
        <div className={styles.register_link}>
          Don&apos;t have an account?
          <Link href="/register"> Register here!</Link>
        </div>
        <button
          disabled={progress}
          type="submit"
          className={styles.submit_login_button}
        >
          Sign In
        </button>
      </div>
    </form>
  );
}
