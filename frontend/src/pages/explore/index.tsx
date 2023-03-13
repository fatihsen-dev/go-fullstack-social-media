import Navbar from "@/components/Navbar/Navbar";
import Head from "next/head";
import styles from "./Exxplore.module.scss";

const Post = () => {
   return (
      <>
         <Head>
            <title>Explore</title>
            <meta name='description' content='Generated by create next app' />
            <meta name='viewport' content='width=device-width, initial-scale=1' />
            <link rel='icon' href='/favicon.ico' />
         </Head>
         <Navbar />
         <div className={styles.Container}>Explore Page</div>
      </>
   );
};

export default Post;