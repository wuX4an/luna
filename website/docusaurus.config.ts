import path from "path";
import { themes as prismThemes } from "prism-react-renderer";
import type { Config } from "@docusaurus/types";
import type * as Preset from "@docusaurus/preset-classic";

const config: Config = {
  title: "Luna",
  tagline: "Effortless Lua scripting for modern applications",
  favicon: "img/favicon.ico",
  future: {
    v4: true,
  },
  url: "https://wux4an.github.io",
  baseUrl: "/luna",
  organizationName: "wux4an",
  projectName: "luna",
  onBrokenLinks: "throw",
  onBrokenMarkdownLinks: "warn",
  i18n: { defaultLocale: "en", locales: ["en"] },
  presets: [
    [
      "classic",
      {
        docs: {
          sidebarPath: "./sidebars.ts",
          path: path.resolve(__dirname, "../docs"), // <-- Aquí apuntas a la carpeta real
          editUrl: "https://github.com/wux4an/luna/tree/main/docs",
        },
        theme: { customCss: "./src/css/custom.css" },
      } satisfies Preset.Options,
    ],
  ],
  themeConfig: {
    image: "img/luna.png",
    docs: {
      sidebar: { hideable: true, autoCollapseCategories: true },
    },
    navbar: {
      title: "Luna",
      logo: { alt: "Luna Logo", src: "img/logo.png" },
      items: [
        {
          type: "docSidebar",
          sidebarId: "tutorialSidebar",
          position: "left",
          label: "Tutorial",
        },
        {
          href: "https://github.com/wux4an/luna",
          label: "GitHub",
          position: "right",
        },
      ],
    },
    footer: {
      style: "dark",
      links: [
        { title: "Docs", items: [{ label: "Tutorial", to: "/docs" }] },
        {
          title: "Community",
          items: [
            {
              label: "Github Discussions",
              href: "https://github.com/wux4an/luna/discussions",
            },
          ],
        },
        {
          title: "More",
          items: [{ label: "GitHub", href: "https://github.com/wux4an/luna" }],
        },
      ],
      copyright: `Copyright © ${new Date().getFullYear()} Luna. Built with ❤️ and Lua`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
      additionalLanguages: ["lua", "toml"],
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
