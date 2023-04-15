import React, { useEffect, useState } from "react";
import { ArrowSmUpIcon } from "@heroicons/react/solid";

const GoToTop = () => {
  const [isVisible, setIsVisible] = useState(false);

  const goToBtn = () => {
    window.scrollTo({ top: 0, left: 0, behavior: "smooth" });
  };

  const listenToScroll = () => {
    let heightToHidden = 20;
    const winScroll =
      document.body.scrollTop || document.documentElement.scrollTop;

    if (winScroll > heightToHidden) {
      setIsVisible(true);
    } else {
      setIsVisible(false);
    }
  };

  useEffect(() => {
    window.addEventListener("scroll", listenToScroll);
    return () => window.removeEventListener("scroll", listenToScroll);
  }, []);

  return (
    <>
      {isVisible && (
        <div className="font-semibold text-4xl w-12 h-12 cursor-pointer bg-white text-black dark:bg-gray-900 dark:text-white  shadow-md rounded-full position: fixed bottom-5 right-5 z-50 flex justify-center items-center" onClick={goToBtn}>
         <ArrowSmUpIcon className="h-6 w-6 text-gray-500" />
        </div>
      )}
    </>
  );
};
export default GoToTop;