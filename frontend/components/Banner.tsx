import Image from "next/image";
import { useRouter } from "next/router";
const Banner = () => {
  const router = useRouter();
  return (
    <div className="relative h-[300px] sm:h-[400px] lg:h-[600px] xl:h-[700px] 2xl:h-[800px]">
      <Image
        alt="banner"
        src="https://images.unsplash.com/photo-1501612780327-45045538702b?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1170&q=80"
        layout="fill"
        objectFit="cover"
      />

      <div className="   absolute top-1/3 w-full text-center">
        <p className="text-sm sm:text-lg font-semibold text-white ">
          Don't know what to listen? Alright.
        </p>
        <button
          onClick={() => {
            router.push("/songs");
          }}
          className="text-purple-700 font-bold bg-white px-8 py-4 my-3 rounded-full shadow-md hover:shadow-2xl active:scale-90 transition duration-150"
        >
          Discover
        </button>
      </div>
    </div>
  );
};

export default Banner;
