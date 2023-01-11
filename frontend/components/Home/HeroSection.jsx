import styles from "./Styles/HeroSection.module.css";
import Image from "next/image";

export default function HeroSection() {
  return (
    <section className={`${styles.hero} py-5 py-md-4`}>
      <div className="container px-5">
        <div className="row px-lg-5">
          <div className="col-md-6 d-flex justify-content-center align-items-center order-2 order-md-1">
            <div className="d-flex flex-column" style={{ gap: ".625em" }}>
              <h1 className={styles.hero_title}>Amet minim mollit non.</h1>
              <h2 className={styles.hero_subtitle}>
                Amet minim mollit non deserunt ullamco est sit aliqua dolor do
                Amet minim mollit non
              </h2>
              <div className={styles.hero_info}>
                <div className="d-flex" style={{ gap: "1.25em" }}>
                  <div className={styles.hero_info_icon}>
                    <svg
                      className={styles.hero_info_svg_icon}
                      xmlns="http://www.w3.org/2000/svg"
                      viewBox="0 0 24 24"
                      style={{
                        fill: "#82AAE3",
                      }}
                    >
                      <path d="M20.56,3.34a1,1,0,0,0-1-.08l-17,8a1,1,0,0,0-.57.92,1,1,0,0,0,.6.9L8,15.45v6.72L13.84,18l4.76,2.08a.93.93,0,0,0,.4.09,1,1,0,0,0,.52-.15,1,1,0,0,0,.48-.79l1-15A1,1,0,0,0,20.56,3.34ZM18.1,17.68l-5.27-2.31L16,9.17,8.35,13.42,5.42,12.13,18.89,5.79Z" />
                    </svg>
                  </div>
                  <div className="d-flex flex-column">
                    <div className={styles.hero_info_title}>5M +</div>
                    <div className={styles.hero_info_description}>
                      Downloaded
                    </div>
                  </div>
                </div>
                <div className="d-flex" style={{ gap: "1.25em" }}>
                  <div className={styles.hero_info_icon}>
                    <svg
                      className={styles.hero_info_svg_icon}
                      fill="none"
                      viewBox="0 0 24 24"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <path
                        d="M4 13.999 13 14a2 2 0 0 1 1.995 1.85L15 16v1.5C14.999 21 11.284 22 8.5 22c-2.722 0-6.335-.956-6.495-4.27L2 17.5v-1.501c0-1.054.816-1.918 1.85-1.995L4 14ZM15.22 14H20c1.054 0 1.918.816 1.994 1.85L22 16v1c-.001 3.062-2.858 4-5 4a7.16 7.16 0 0 1-2.14-.322c.336-.386.607-.827.802-1.327A6.19 6.19 0 0 0 17 19.5l.267-.006c.985-.043 3.086-.363 3.226-2.289L20.5 17v-1a.501.501 0 0 0-.41-.492L20 15.5h-4.051a2.957 2.957 0 0 0-.595-1.34L15.22 14H20h-4.78ZM4 15.499l-.1.01a.51.51 0 0 0-.254.136.506.506 0 0 0-.136.253l-.01.101V17.5c0 1.009.45 1.722 1.417 2.242.826.445 2.003.714 3.266.753l.317.005.317-.005c1.263-.039 2.439-.308 3.266-.753.906-.488 1.359-1.145 1.412-2.057l.005-.186V16a.501.501 0 0 0-.41-.492L13 15.5l-9-.001ZM8.5 3a4.5 4.5 0 1 1 0 9 4.5 4.5 0 0 1 0-9Zm9 2a3.5 3.5 0 1 1 0 7 3.5 3.5 0 0 1 0-7Zm-9-.5c-1.654 0-3 1.346-3 3s1.346 3 3 3 3-1.346 3-3-1.346-3-3-3Zm9 2c-1.103 0-2 .897-2 2s.897 2 2 2 2-.897 2-2-.897-2-2-2Z"
                        fill="#82aae3"
                      />
                    </svg>
                  </div>
                  <div className="d-flex flex-column">
                    <div className={styles.hero_info_title}>1M +</div>
                    <div className={styles.hero_info_description}>
                      User&apos;s Active
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div className="col-md-6 d-flex justify-content-center align-items-center order-1 order-md-2 mb-3 mb-md-0">
            <div className={styles.image_container}>
              <Image
                fill
                src="/Assets/hero-image.png"
                alt="hero-image"
                sizes="contain"
                priority
              ></Image>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
