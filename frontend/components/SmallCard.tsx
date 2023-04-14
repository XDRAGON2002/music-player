import Image from "next/image";
import { ISong } from "@/types";
import { HeartIcon } from "@heroicons/react/solid";

const SmallCard = (songData: ISong) => {
  return (
    // <div className="flex items-center m-2 mt-5 space-x-4 rounded-xl cursor-pointer hover:bg-gray-100 hover:scale-105 transition transform duration-200 ease-out">
    //   <div className="relative  h-16 w-16">
    //     <Image alt = "image "src={songData.image} layout="fill" className="rounded-lg" />
    //   </div>
    //   <div>
    //     <h2 className="font-bold">{songData.songname}</h2>
    //     {/* <h3 className="text-gray-500 font-semibold">{distance}</h3> */}
    //   </div>
    // </div>

    <div className="dark:bg-gray-900 flex flex-col items-center bg-white border border-gray-200 rounded-lg shadow-lg md:flex-row md:max-w-xl hover:bg-gray-100 hover:scale-105 transition transform duration-200 ease-out dark:border-0 ">
      <img
        className="object-cover w-full rounded-t-lg h-96 md:h-auto md:w-48 md:rounded-none md:rounded-l-lg"
        src={songData.image}
        alt=""
      />
      <div className="flex flex-col justify-between p-4 leading-normal">
        <h5 className="mb-2 text-md font-bold tracking-tight text-gray-900 dark:text-white">
          {songData.songname}
        </h5>
        <p className="mb-3 text-sm font-normal text-gray-700 dark:text-gray-400">
          {songData.songartists}
        </p>
        <div className="flex">
          <div>
            <HeartIcon className="text-right  mt-[4.5] h-6 pr-3 pb-3" />
          </div>
          <p className="  mt-3 text-xs font-normal text-gray-700 dark:text-gray-400">
            {songData.likes}
          </p>
        </div>
      </div>
    </div>
  );
};

export default SmallCard;
