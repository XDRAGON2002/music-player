import SmallCard from "@/components/SmallCard";
import React, { useEffect, useState } from "react";
import { ISong } from "@/types";
import Header from "@/components/Header";
import Link from "next/link";
import axios from "axios";
import { InferGetServerSidePropsType } from "next";
import InfiniteScroll from "react-infinite-scroll-component";
import { useRouter } from "next/router";
import GoToTop from "@/components/GoToTop";

const home = ({
  data,
}: InferGetServerSidePropsType<typeof getServerSideProps>) => {
  const [songs, setSongs] = useState<ISong[]>(data);
  const [hasMore, setHasMore] = useState(true);
  const [length, setLength] = useState(data.length)
  const [pages, setPages] = useState(1);
  const router = useRouter();

  const getMoreSongs = async () => {
    setPages(pages+1)
    const res = await axios.get(
      `http://localhost:5000/api/song/page/${pages}`
    );
    const newSongs = await res.data;
    setLength(newSongs.length)
    console.log(newSongs)
    setSongs((songs) => [...songs, ...newSongs]);
  };

  useEffect(() => {
    setHasMore(length > 0 ? true : false);
    console.log(hasMore)
  }, [songs]);

  return (
    <>
      <Header />
      <main className=" max-w-7xl mx-auto px-10 pt-4 ">
        <section className="pt-6">
          <h2 className="text-4xl font-semibold pb-8">Explore Songs</h2>
          <hr />
          <br />
          <br />

          {/* pulling data from a file */}

          
            <InfiniteScroll
              dataLength={songs.length}
              next={getMoreSongs}
              hasMore={true}
              loader={<h4>Loading...</h4>}
              endMessage={
                <p style={{ textAlign: "center" }}>
                  <b>Yay! You have seen it all</b>
                </p>
              }
            >
              <div className="pt-4 px-4  gap-[2.75rem] grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-3 ">
              {songs?.length > 0 &&
                songs.map((item) => (
                  <Link target="_blank" href={`/songs/${item._id}`}>
                    <SmallCard
                      _id={item._id}
                      image={item.image}
                      songname={item.songname}
                      songartists={item.songartists}
                      likes={item.likes}
                    />
                  </Link>
                ))}

              </div>


              {songs.length == 0 && (
                <div className="flex  w-full  justify-center gap-2 text-gray-400">
                  <span className="h-6 w-6 mt-60 block rounded-full border-4 dark:border-white border-black animate-spin"></span>
                </div>
              )}
              <GoToTop/>
            </InfiniteScroll>
            
        </section>
      </main>

      <br />
    </>
  );
};

export async function getServerSideProps() {
  // Fetch the data for the current page
  const res = await axios.get(`http://localhost:5000/api/song/page/1`);
  const data = await res.data;
  return {
    props: {
      data: data,
    },
  };
}


export default home;
