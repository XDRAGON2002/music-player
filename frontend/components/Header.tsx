import Image from "next/image";
// @ts-ignore
import {
  SearchIcon,
  GlobeAltIcon,
  UserCircleIcon,
  MenuIcon,
  UsersIcon,
} from "@heroicons/react/solid";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";
import mypic from "../assets/images/audiofi_logo.png";
import { useTheme } from "next-themes";
import { SunIcon, MoonIcon } from "@heroicons/react/solid";
import Link from "next/link";

const Header = () => {
  const { systemTheme, theme, setTheme } = useTheme();

  const router = useRouter();
  const [searchInput, setSearchInput] = useState("");

  const [mounted, setMounted] = useState(false);

  useEffect(() => {
    setMounted(true);
  }, []);

  const renderThemeChanger = () => {
    if (!mounted) return null;

    const currentTheme = theme === "system" ? systemTheme : theme;

    if (currentTheme === "dark") {
      return (
        <SunIcon
          className="w-10 h-6   text-yellow-500 "
          role="button"
          onClick={() => setTheme("light")}
        />
      );
    } else {
      return (
        <MoonIcon
          className="w-10 h-6 text-gray-900 "
          role="button"
          onClick={() => setTheme("dark")}
        />
      );
    }
  };

  return (
    <header className="dark:bg-gray-900 w-full drop-shadow-sm sticky top-0 z-50 bg-white shadow-md grid grid-cols-3 py-3 px-2 md:px-2">
      {/* left */}
      <div
        onClick={() => {
          router.push("/");
        }}
        className="relative flex items-center h-15 cursor-pointer w-[60px]"
      >
        <Image alt="image" src={mypic} width={60} height={60} />
        <div className="invisible sm:visible  font-bold  font-righteous text-lg">
          AudioFi
        </div>
      </div>

      {/* middle  search box*/}
      <div className="flex items-center w-full rounded-full border-1 h-[40px] mt-1 shadow-md">
        <input
          value={searchInput}
          onChange={(e) => setSearchInput(e.target.value)}
          className="pl-5 bg-transparent outline-none flex-grow text-gray-600 placeholder:invisible sm:placeholder:visible"
          type="text"
          placeholder="start searching"
        />
        <SearchIcon className=" hidden md:mx-2  md:inline-flex h-8  rounded-full p-2 cursor-pointer" />
      </div>

      {/* right */}
      <div className="flex space-x-4 items-center justify-end text-black">
        <div className="flex   font-semibold text-base items-center space-x-2  p-2 rounded-full">
          {/* <UserCircleIcon className='dark:text-white h-6 cursor-pointer data-dropdown-toggle="dropdown"'/> */}

          <button
            id="dropdownDefaultButton"
            data-dropdown-toggle="dropdown"
            className="text-black font-medium   text-[0.55rem]  text-center inline-flex items-center  dark:text-white "
            type="button"
          >
            <svg
              fill="none"
              stroke="currentColor"
              strokeWidth={1.5}
              width={27}
              height={27}
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
              aria-hidden="true"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                d="M17.982 18.725A7.488 7.488 0 0012 15.75a7.488 7.488 0 00-5.982 2.975m11.963 0a9 9 0 10-11.963 0m11.963 0A8.966 8.966 0 0112 21a8.966 8.966 0 01-5.982-2.275M15 9.75a3 3 0 11-6 0 3 3 0 016 0z"
              />
            </svg>
          </button>
          {/* <!-- Dropdown menu --> */}
          <div
            id="dropdown"
            className="z-10 hidden bg-white divide-y divide-gray-100 rounded-lg shadow w-44 dark:bg-gray-700"
          >
            <ul
              className="py-2 text-sm text-gray-700 dark:text-gray-200"
              aria-labelledby="dropdownDefaultButton"
            >
              <li>
                <Link
                  href="#"
                  className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                >
                  Dashboard
                </Link>
              </li>
              <li>
                <Link
                  href="#"
                  className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                >
                  Settings
                </Link>
              </li>
              <li>
                <Link
                  href="#"
                  className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                >
                  Earnings
                </Link>
              </li>
              <li>
                <Link
                  href="#"
                  className="block px-4 py-2 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white"
                >
                  Sign out
                </Link>
              </li>
            </ul>
          </div>

          {renderThemeChanger()}
        </div>
      </div>
    </header>
  );
};

export default Header;
