import Image from "next/image";
import { ISong } from "@/types";

const SmallCard = (songData: ISong) => {
  return (
    <div className="flex items-center m-2 mt-5 space-x-4 rounded-xl cursor-pointer hover:bg-gray-100 hover:scale-105 transition transform duration-200 ease-out">
      <div className="relative  h-16 w-16">
        <Image alt = "image "src={songData.image} layout="fill" className="rounded-lg" />
      </div>
      <div>
        <h2 className="font-bold">{songData.songName}</h2>
        {/* <h3 className="text-gray-500 font-semibold">{distance}</h3> */}
      </div>
    </div>
  );
};

export default SmallCard;