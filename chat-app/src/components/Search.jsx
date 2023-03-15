import React from "react";

const Search = () => {
    return (
        <div className="search">
            <div className="searchForm">
                <input type="text"  placeholder="Find user"/>
            </div>
            <div className="userChat">
                <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQArTzo_6ODTO8yjuSqaWUEAbvVM3oFubu5VWrRnv9Kh4pgWRPNsNUU--p9lOrQYmElzR0&usqp=CAU" alt="" />
                <div className="userChatInfo">
                    <span>chesly</span>
                </div>
            </div>
        </div>
    )
}
export default Search
