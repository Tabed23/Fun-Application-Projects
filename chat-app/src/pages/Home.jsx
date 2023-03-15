import React from "react";
import Sidebar from "../components/Sidebar";
import Chat from "../components/Chat";

const Home = () => {
    return (
        <div className="home">
            <dev className='container'>
                <Sidebar />
                <Chat />

            </dev>
        </div>
    )
}
export default Home
