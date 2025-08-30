import type { ReactNode } from "react";
import clsx from "clsx";
import Heading from "@theme/Heading";
import styles from "./styles.module.css";

type FeatureItem = {
  title: string;
  description: ReactNode;
};

const FeatureList: FeatureItem[] = [
  {
    title: "Lightweight & Fast",
    description: (
      <>
        Luna is a minimal, high-performance Lua runtime designed to execute
        scripts efficiently, with zero bloat and lightning-fast startup.
      </>
    ),
  },
  {
    title: "Modern Std Library",
    description: (
      <>
        Access a rich set of modern standard libraries for HTTP, async, crypto,
        SQLite, CLI, templating, and more—ready to use out of the box.
      </>
    ),
  },
  {
    title: "Developer-Friendly CLI",
    description: (
      <>
        Luna provides a powerful, easy-to-use CLI for running, building,
        testing, and documenting your Lua projects with intuitive commands.
      </>
    ),
  },
  {
    title: "Extensible & Modular",
    description: (
      <>
        Build your own modules, extend the runtime, or embed resources inside
        your projects. Luna is fully modular and adapts to your workflow.
      </>
    ),
  },
  {
    title: "Cross-Platform Ready",
    description: (
      <>
        Compile your scripts into standalone executables for Linux, macOS, or
        Windows without extra toolchains.
      </>
    ),
  },
  {
    title: "Interactive REPL & Watch Mode",
    description: (
      <>
        Experiment in real time with an interactive REPL, hot reload your
        scripts, and get immediate feedback—perfect for rapid development.
      </>
    ),
  },
];

function Feature({ title, description }: FeatureItem) {
  return (
    <div className={clsx("col col--4")}>
      <div className="text--center"></div>
      <div className="text--center padding-horiz--md">
        <Heading as="h3">{title}</Heading>
        <p>{description}</p>
      </div>
    </div>
  );
}

export default function HomepageFeatures(): ReactNode {
  return (
    <section className={styles.features}>
      <div className="container">
        <div className="row">
          {FeatureList.map((props, idx) => (
            <Feature key={idx} {...props} />
          ))}
        </div>
      </div>
    </section>
  );
}
