import {useState, useEffect} from 'react';
import {CreateSongRequest, Song} from "./generated/service_pb"
import {MusicPlayerClient} from "./generated/ServiceServiceClientPb"

const client = new MusicPlayerClient("http://localhost:8080")

function App() {
    const [song, setSong] = useState<string | undefined>();
    const getTemp = () => {
        console.log("called")

        const createSongRequest = new CreateSongRequest();
        const song = new Song();
        song.setSongtitle('Test test');
        song.setElapsed(1000);
        createSongRequest.setSong(song);
        const resp = client.createSong(createSongRequest,{});
        resp.then((e) => {setSong("ok")});
    };


    useEffect(()=>{
        getTemp()
    },[]);

    return (
        <div>
            <p className="w-full text-center pt-48 text-5xl text-blue-500">
                Hello {song}
            </p>
        </div>
    );
}

export default App;
