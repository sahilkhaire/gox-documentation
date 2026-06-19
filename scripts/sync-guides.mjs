import { cpSync, existsSync, mkdirSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

const root = join(dirname(fileURLToPath(import.meta.url)), "..");
const goxDocs = join(root, "..", "gox", "docs");

const copies = [
  [join(goxDocs, "architecture.md"), join(root, "docs", "guide", "architecture.md")],
  [join(goxDocs, "CHEATSHEET.md"), join(root, "docs", "reference", "cheatsheet.md")],
];

const migrationSrc = join(goxDocs, "migration");
const migrationDst = join(root, "docs", "migration");

if (!existsSync(goxDocs)) {
  console.warn("gox/docs not found at", goxDocs, "— skipping sync");
  process.exit(0);
}

for (const [src, dst] of copies) {
  mkdirSync(dirname(dst), { recursive: true });
  cpSync(src, dst);
  console.log("copied", src, "→", dst);
}

mkdirSync(migrationDst, { recursive: true });
cpSync(migrationSrc, migrationDst, { recursive: true });
console.log("copied migration guides →", migrationDst);
