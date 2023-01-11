import { Blocks } from "react-loader-spinner";

export default function Loading() {
  return (
    <div
      className="position-absolute d-flex justify-content-center align-items-center bg-white"
      style={{
        inset: "0",
        background: "rgba(255,255,255,0.6)",
        zIndex: "9999",
      }}
    >
      <Blocks
        visible={true}
        height="80"
        width="80"
        ariaLabel="blocks-loading"
        wrapperStyle={{}}
        wrapperClass="blocks-wrapper"
      />
    </div>
  );
}
