import React from "react";
import "./App.css";
import { useQuery } from "react-query";
import { apiClientV1 } from "./api/apiClient";

type User = {
   Id: number;
  FirstName: string;
  LastName: string;
};
type Users = User[];

function App() {
  let a = 10;

  const { data, isLoading } = useQuery<Users>("get/users", async () => {
    const { data } = await apiClientV1.get("/users");
    return data;
  });

  return (
    <div className="App">
      <h1>Hello</h1>
      {isLoading ? (
        "loading...5"
      ) : (
        <ul>
          {data?.map((d: User) => {
            return (
              <li key={d.Id}>
                {d.Id}, {d.FirstName}, {d.LastName}
              </li>
            );
          })}
        </ul>
      )}
    </div>
  );
}

export default App;
