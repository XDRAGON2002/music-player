import Image from 'next/image'
// @ts-ignore
import {SearchIcon, GlobeAltIcon, UserCircleIcon, MenuIcon, UsersIcon} from '@heroicons/react/solid'
import { useRouter } from 'next/router';
import { useState } from 'react';
import mypic from '../assets/images/audiofi_logo.png'

const Header = () => {
    const router = useRouter();
    const [searchInput, setSearchInput] = useState("");


    return <header className='sticky top-0 z-50 bg-white shadow-md grid grid-cols-3 py-3 px-2 md:px-2'>
  {/* left */}
  <div onClick={()=>{router.push("/")}} className='relative flex items-center h-15 cursor-pointer w-[60px]'>
    <Image alt='image' src={mypic}
   width={60} 
   height={60}/>
   <div className='font-bold text-black font-righteous text-lg'>AudioFi</div>
  </div> 

  {/* middle  search box*/}
  <div className='flex items-center rounded-full border-1 h-[40px] mt-1 shadow-md'>
    <input 
    value={searchInput}
    onChange ={(e)=>setSearchInput(e.target.value)}
     className='pl-5 bg-transparent outline-none flex-grow text-gray-600' type="text" placeholder='start searching'/>
    <SearchIcon className=" hidden md:mx-2  md:inline-flex h-8 bg-black text-white rounded-full p-2 cursor-pointer"/>
  </div>

  {/* right */}
  <div className='flex space-x-4 items-center justify-end text-black'>
    <div className='flex    text-black font-semibold text-base items-center space-x-2 border-2 p-2 rounded-full'>
        <UserCircleIcon className='h-6 cursor-pointer'/>
    </div>
  </div>

  </header>

} 

export default Header