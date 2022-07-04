import { ApplicationData } from "@/types/apps";
import { getLangIcon } from "@/utils/icons";
import Link from "next/link";
import { useState } from "react";

export default function EditAppCard({
  id,
  name,
  namespace,
  languages,
  kind,
}: ApplicationData) {
  return (
    <div className="shadow-lg border border-gray-200 rounded-lg bg-white hover:bg-gray-100 cursor-pointer">
      <Link href={`/source/edit/${name}`}>
        <a className="flex flex-row p-3 items-center space-x-4">
          {getLangIcon(languages[0], "w-12 h-12")}
          <div className="flex flex-col items-start">
            <div className="font-bold">{name}</div>
            <div>{kind}</div>
            <div>namespace: {namespace}</div>
          </div>
        </a>
      </Link>
    </div>
  );
}
