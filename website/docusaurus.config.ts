import path from "path";
import { themes as prismThemes } from "prism-react-renderer";
import type { Config } from "@docusaurus/types";
import type * as Preset from "@docusaurus/preset-classic";

const config: Config = {
  title: "Luna",
  tagline:
    "Lua runtime with extended standard library, ready to build for common devices",
  favicon: "img/luna.svg",
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
          path: path.resolve(__dirname, "../docs"),
          editUrl: "https://github.com/wux4an/luna/tree/main/docs",
        },
        theme: { customCss: "./src/css/custom.css" },
      } satisfies Preset.Options,
    ],
  ],
  plugins: [
    [
      require.resolve("@easyops-cn/docusaurus-search-local"),
      {
        hashed: true,
        language: ["en"],
        docsRouteBasePath: "/",
        indexBlog: false,
        highlightSearchTermsOnTargetPage: true,
        explicitSearchResultPath: true,
        docsDir: "../docs",
      },
    ],
  ],
  themeConfig: {
    image: "img/luna.svg",
    docs: {
      sidebar: { hideable: false, autoCollapseCategories: true },
    },
    colorMode: {
      disableSwitch: true,
    },
    navbar: {
      title: "Luna",
      logo: { alt: "Luna Logo", src: "img/luna.svg" },
      items: [
        {
          type: "search",
          position: "left",
        },
        {
          type: "docSidebar",
          sidebarId: "tutorialSidebar",
          position: "right",
          label: "Docs",
        },
        {
          type: "html",
          position: "right",
          value: '<span style="color:#8d8e94;">|</span>',
        },
        {
          href: "https://github.com/wux4an/luna",
          label: "GitHub",
          position: "right",
        },
        {
          type: "html",
          position: "right",
          value: '<span style="margin:0 2px;"></span>',
        },
      ],
    },
    prism: {
      theme: prismThemes.dracula,
      darkTheme: prismThemes.dracula,
      additionalLanguages: ["lua", "toml"],
    },
  } satisfies Preset.ThemeConfig,
};

export default config;
