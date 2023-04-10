import axios from "axios";
import { useEffect, useState } from "react";

export const Songs = async () => {
    // const [SongData, setSongData] = useState<any[]>([]);
    // useEffect(() => {
    //     axios.get('https://localhost:5000/songs/').then(function(response) {
    //         setSongData(response.data);
    //         console.log(SongData)
    //         }).catch(function(error) {
    //         console.log(error);
    //     });
    //     },[])
    // console.log(SongData)
    // return SongData

    try{
        const response = await axios.get('http://localhost:5000/api/song/')
        return {success: true, data: response.data}
    }catch(err:any){
        console.log("Error in fetching songs: ", err.message)
        return {succes: false, message: err.message}
    }
}

