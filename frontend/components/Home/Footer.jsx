import styles from "./Styles/Footer.module.css";

export default function Footer() {
  return (
    <footer className={`text-center text-lg-start ${styles.section}`}>
      <div className="container p-4 pb-0">
        <section>
          <div className="row mb-5">
            <div className="col-md-3 col-lg-3 col-xl-3 mx-auto mt-3">
              <h6 className={`mb-4 ${styles.logo}`}>
                My
                <br />
                Gallery.
              </h6>
            </div>
            <hr className="w-100 clearfix d-md-none" />
            <div className="col-md-3 col-lg-2 col-xl-2 mx-auto mt-3">
              <h6 className={`mb-4 ${styles.col_header}`}>Lorem Ipsum</h6>
              <p>
                <a className={styles.col_text}>Lorem</a>
              </p>
              <p>
                <a className={styles.col_text}>Lorem</a>
              </p>
              <p>
                <a className={styles.col_text}>Lorem</a>
              </p>
            </div>
            <hr className="w-100 clearfix d-md-none" />
            <div className="col-md-3 col-lg-2 col-xl-2 mx-auto mt-3">
              <h6 className={`mb-4 ${styles.col_header}`}>Lorem Ipsum</h6>
              <p>
                <a className={styles.col_text}>Lorem</a>
              </p>
              <p>
                <a className={styles.col_text}>Lorem</a>
              </p>
              <p>
                <a className={styles.col_text}>Lorem</a>
              </p>
            </div>
            <hr className="w-100 clearfix d-md-none" />
            <div className="col-md-3 col-lg-2 col-xl-2 mx-auto mt-3">
              <h6 className={`mb-4 ${styles.col_header}`}>Lorem Ipsum</h6>
              <p>
                <a className={styles.col_text}>Lorem</a>
              </p>
              <p>
                <a className={styles.col_text}>Lorem</a>
              </p>
              <p>
                <a className={styles.col_text}>Lorem</a>
              </p>
            </div>
          </div>
        </section>
        <hr className="my-3" />
        <section className="p-3 pt-0">
          <div className="row d-flex align-items-center">
            <div className="col-md-7 col-lg-8 text-center text-md-start">
              <div className={`p-3 ${styles.copyright}`}>
                Copyright © 2022 MyGallery. All rights reserved.
              </div>
            </div>
            <div className="col-md-5 col-lg-4 ml-lg-0 text-center text-md-end"></div>
          </div>
        </section>
      </div>
    </footer>
  );
}
