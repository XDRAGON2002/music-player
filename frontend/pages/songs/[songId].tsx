import { useRouter } from 'next/router'

const PlaySong = () => {
  const router = useRouter()
  const { songId } = router.query  

  return <p>PlaySong: {songId}</p>
}

export default PlaySong