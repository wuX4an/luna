import type { ReactNode } from "react";
import clsx from "clsx";
import Link from "@docusaurus/Link";

type ButtonProps = {
  children: ReactNode;
  to: string; // ahora usamos "to" como prop estándar de Link
  className?: string;
};

export default function Button({ children, to, className }: ButtonProps) {
  return (
    <Link
      className={clsx(className)}
      to={to} // Link maneja la navegación internamente
      style={{ textDecoration: "none" }}
    >
      {children}
    </Link>
  );
}
