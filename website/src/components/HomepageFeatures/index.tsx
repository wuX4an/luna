import type { ReactNode } from "react";
import { FiExternalLink } from "react-icons/fi";
import { IoIosWarning } from "react-icons/io";
import clsx from "clsx";
import Heading from "@theme/Heading";
import styles from "./styles.module.css";
import Button from "@site/src/components/ui/button";

type FeatureItem = {
  title: string;
  description: ReactNode;
};

const FeatureList: FeatureItem[] = [
  {
    title: "Install",
    description: (
      <>
        Install Luna by downloading the <code>nightly</code> build from the
        official releases.
        <br />
        <br />
        Nightly builds are updated frequently with the latest features and
        improvements
        <br />
        <br />
        <Button
          className={styles.buttonDownload}
          to="https://github.com/wuX4an/luna/releases"
        >
          Go To Github Releases
          <FiExternalLink
            size={15}
            style={{
              marginLeft: "0.3rem",
              position: "relative",
              top: "1px",
            }}
          />
        </Button>
        <br />
        <br />
        <p className="warning-block">
          <IoIosWarning
            size={18}
            style={{
              marginRight: "0.5rem",
              position: "relative",
              top: "2px",
            }}
          />
          Warning: This project is currently in an experimental stage.
        </p>
      </>
    ),
  },
  {
    title: "Initialize",
    description: (
      <>
        Initialize a new project with <code>luna init</code>, which sets up the
        necessary structure to get you started quickly.
        <br />
        <img src="/luna/demos/init.gif" alt="Init demo" width={300} />
      </>
    ),
  },
  {
    title: "Build",
    description: (
      <>
        Build your scripts for any device with
        <br />
        <code>luna build</code>. This bundles your Lua code into standalone
        runtime ready to run.
        <img src="/luna/demos/build.gif" alt="Build demo" width={300} />
      </>
    ),
  },
];

function Feature({ title, description }: FeatureItem) {
  return (
    <div className={styles.featureBox}>
      <Heading as="h3">{title}</Heading>
      <p>{description}</p>
    </div>
  );
}

export default function HomepageFeatures(): ReactNode {
  return (
    <section className={styles.features}>
      <div className={styles.row}>
        {FeatureList.map((props, idx) => (
          <Feature key={idx} {...props} />
        ))}
      </div>
    </section>
  );
}
