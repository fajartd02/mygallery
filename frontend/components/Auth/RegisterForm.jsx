import Link from "next/link";
import { useState } from "react";
import { postJsonOpt } from "../../helper/options";
import { toast } from "react-toastify";
import styles from "./Styles/LoginForm.module.css";
import { capitalize } from "../../helper/utils";

export default function RegisterForm() {
  const [progress, setProgress] = useState(false);

  async function submitHandler(e) {
    e.preventDefault();
    setProgress(true);

    const data = {
      fullName: e.target.fullname.value,
      email: e.target.email.value,
      password: e.target.password.value,
    };

    try {
      const res = await fetch(
        `${process.env.NEXT_PUBLIC_URL}/register`,
        postJsonOpt(data)
      );

      const resData = await res.json();
      const success = resData?.status && true;

      if (success) {
        toast.success(`${resData.status} 👌!`);
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
    <>
      <form
        onSubmit={submitHandler}
        className="d-flex justify-content-center align-items-center"
      >
        <div className={styles.form_container}>
          {progress && <div className={styles.overlay}></div>}
          <h1 className={styles.form_title}>Register</h1>
          <div className={styles.form_input_container}>
            <label htmlFor="fullname">Name</label>
            <input
              id="fullname"
              type="text"
              placeholder="Fullname..."
              name="fullname"
              required
            />
          </div>
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
            Already have account? <Link href="/login">Login here!</Link>
          </div>
          <button
            disabled={progress}
            type="submit"
            className={styles.submit_login_button}
          >
            Register
          </button>
        </div>
      </form>
    </>
  );
}
