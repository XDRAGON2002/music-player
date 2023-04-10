import SmallCard from '@/components/SmallCard'
import {Songs} from '@/fetchdata/GetSongs'
import React, { useEffect, useState } from 'react'
import { ISong } from '@/types'
import Header from '@/components/Header'

const home = () => {

    const [loading, setLoading] = useState(true);

    const [songs, setSongs] = useState<ISong[]>([])
    useEffect(() => {
        (async () => {
            setLoading(true);
            const _songs = await Songs()
            if(!_songs.success) console.log("Failed to fetch songs")

            setSongs(_songs.data)
            setLoading(false);
        })()
    },[])

  return (
    <>
    <Header/>
    <main className='max-w-7xl mx-auto px-8 '>
        <section className='pt-6'>
          <h2 className='text-4xl font-semibold pb-5'>Explore Songs</h2>

          {/* pulling data from a file */}

        <div className='grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4'>
          {songs?.length > 0 && songs.map((item)=>(
           <SmallCard 
            id={item.id}
            image={item.image} 
            songName={item.songName} 
            />
          ))}
        </div>
        {songs.length == 0 &&
        <div className="flex  w-full  justify-center gap-2 text-white">
            <span className="h-6 w-6 mt-60 block rounded-full border-4 border-black animate-spin"></span>
                  
        </div>} 
         
        </section>
        </main>

    </>
  )
}

export default home