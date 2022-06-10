import Head from "next/head";
import styles from "../styles/Home.module.css";
import React from "react";
import { TextField } from "@mui/material";
import NextLink from "next/link";
import { Link as MUILink } from "@mui/material";
import { useRouter } from "next/router";
import TopRanking from "../components/top-ranking";

export default function Home() {
  const router = useRouter();

  return (
    <div className={styles.container}>
      <Head>
        <title>Create Next App</title>
        <meta name="description" content="Generated by create next app" />
        <link rel="icon" href="/color-graph.ico" />
      </Head>

      <main className={styles.main}>
        <h1 className={styles.title}>AtCoder Diff Problems</h1>

        <div>
          <TextField
            id="outlined-basic"
            label="user-id"
            variant="outlined"
            fullWidth
            onKeyDown={(event) => {
              if (event.key === "Enter") {
                const userId = (event.target as HTMLInputElement).value;
                if (userId === "") {
                  return;
                }
                router.push({
                  pathname: "users/" + userId,
                });
              }
            }}
          />
        </div>

        <div>
          <TopRanking />
          <NextLink href="/ranking" passHref>
            <MUILink variant="body2">All Ranking</MUILink>
          </NextLink>
        </div>
      </main>

      <footer className={styles.footer}>
        <a href="https://twitter.com/sirogami_main" target="_blank" rel="noopener noreferrer">
          Powered by Sirogamichandayo
        </a>
      </footer>
    </div>
  );
}