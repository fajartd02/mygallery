import Image from "next/image";

export default function MemoryImage({ imageUrl, ...props }) {
  const src = `${process.env.NEXT_PUBLIC_STATIC_IMAGE_URL}${imageUrl}`;
  return (
    <Image
      unoptimized={() => src}
      src={src}
      alt="memory-image"
      {...props}
    ></Image>
  );
}
