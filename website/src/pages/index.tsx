import type { ReactNode } from "react";
import Head from "@docusaurus/Head";
import { FaBook, FaGithub } from "react-icons/fa";
import { PiRocketLaunchFill } from "react-icons/pi";
import { FiExternalLink } from "react-icons/fi";
import { TypeAnimation } from "react-type-animation";
import clsx from "clsx";
import useDocusaurusContext from "@docusaurus/useDocusaurusContext";
import Layout from "@theme/Layout";
import HomepageFeatures from "@site/src/components/HomepageFeatures";
import Button from "@site/src/components/ui/button";
import Heading from "@theme/Heading";
import styles from "./index.module.css";

function HomepageHeader() {
  const { siteConfig } = useDocusaurusContext();
  return (
    <header className={clsx("hero hero--primary", styles.heroBanner)}>
      <div className="container">
        <Heading as="h1" className={clsx("hero__title", styles.heroTitle)}>
          {siteConfig.title}
          <PiRocketLaunchFill
            style={{
              marginLeft: "1.4rem",
              verticalAlign: "middle",
              color: "#06b6d4",
              border: "2px solid #8b5cf6",
              borderRadius: "50%",
              padding: "0.25rem",
              boxShadow: "0 -2px 6px rgba(0,0,0,23)",
              cursor: "pointer",
            }}
            onClick={() =>
              window.open("https://www.nasa.gov/mission/apollo-11/", "_blank")
            }
          />
        </Heading>
        <p className={clsx(styles.heroDesc)}>
          Lua runtime with extended standard library, ready to build for common
          devices
        </p>
        <p className={clsx(styles.heroMotto)}>
          <TypeAnimation
            sequence={[
              "Ready-To-Run and Cross-Compile Lua",
              1500,
              "Zero dependencies required",
              1500,
              "Powered by Go + Lua",
              1500,
              "Built for Devs",
              2000,
            ]}
            wrapper="span"
            speed={40}
            omitDeletionAnimation={true}
            repeat={Infinity}
          />
        </p>
        <div className={clsx(styles.buttonContainer)}>
          <Button className={styles.buttonStart} to="docs">
            <span style={{ display: "flex", alignItems: "center" }}>
              <FaBook
                size={20}
                style={{
                  marginRight: "0.8rem",
                  position: "relative",
                  top: "-1px",
                }}
              />
              Get Started
            </span>
          </Button>
          <Button
            className={styles.buttonGithub}
            to="https://github.com/wux4an/luna"
          >
            <span style={{ display: "flex", alignItems: "center" }}>
              <FaGithub
                size={23}
                style={{
                  marginRight: "0.8rem",
                  position: "relative",
                  top: "-1px",
                }}
              />
              GitHub
              <FiExternalLink
                size={15}
                style={{
                  marginLeft: "0.3rem",
                  position: "relative",
                  top: "-1px",
                }}
              />
            </span>
          </Button>
        </div>
      </div>
    </header>
  );
}

export default function Home(): ReactNode {
  const { siteConfig } = useDocusaurusContext();
  return (
    <Layout
      title={`${siteConfig.title}`}
      description="Lua Toolkit Ready-To-Run"
    >
      <Head>
        {/* Favicon */}
        <link rel="icon" href="/img/luna.svg" />

        {/* Open Graph / Discord / GitHub */}
        <meta property="og:title" content={siteConfig.title} />
        <meta property="og:description" content="Lua Toolkit Ready-To-Run" />
        <meta property="og:image" content="/img/luna.svg" />
        <meta property="og:type" content="website" />
        <meta property="og:url" content={siteConfig.url} />

        {/* Discord rich preview */}
        <meta name="theme-color" content="#06b6d4" />
        <meta property="og:site_name" content={siteConfig.title} />
        <meta property="og:locale" content="en_US" />

        {/* GitHub / general SEO */}
        <meta name="description" content="Lua Toolkit Ready-To-Run" />
        <meta
          name="keywords"
          content="luna, Lua, Go, Cross-compile, Dev Tools, github/wux4an"
        />
        <meta name="author" content="wux4an" />
      </Head>
      <HomepageHeader />
      <main>
        <HomepageFeatures />
      </main>
    </Layout>
  );
}
