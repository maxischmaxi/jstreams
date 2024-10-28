"use client";

import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  getPaginationRowModel,
  PaginationState,
  useReactTable,
} from "@tanstack/react-table";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import {
  Pagination,
  PaginationContent,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
} from "@/components/ui/pagination";
import { useMemo, useState } from "react";
import { ChampionChartEntry } from "../page";
import Image from "next/image";

type Props = {
  data: ChampionChartEntry[];
};

const columnHelper = createColumnHelper<ChampionChartEntry>();
const defaultColumns = [
  columnHelper.accessor("championPoints", {
    header: "Punkte",
  }),
  columnHelper.accessor("name", {
    header: "Name",
    cell: (props) => (
      <div className="flex flex-row flex-nowrap gap-1">
        <Image
          width={48}
          height={48}
          alt={props.row.original.name}
          src={props.row.original.icon}
          className="rounded-full object-cover object-center"
        />
        <p className="text-sm">{props.row.original.name}</p>
      </div>
    ),
  }),
];

export function Champions(props: Props) {
  const [pagination, setPagination] = useState<PaginationState>({
    pageIndex: 0,
    pageSize: 10,
  });

  const pageCount = useMemo(
    () => Math.ceil(props.data.length / pagination.pageSize),
    [pagination.pageSize, props.data.length],
  );

  const table = useReactTable({
    columns: defaultColumns,
    data: props.data.sort((a, b) => b.championLevel - a.championLevel),
    getCoreRowModel: getCoreRowModel(),
    manualPagination: false,
    pageCount,
    getPaginationRowModel: getPaginationRowModel(),
    onPaginationChange: setPagination,
    state: {
      pagination,
    },
  });

  return (
    <>
      <Table>
        <TableCaption>Your Champions</TableCaption>
        <TableHeader>
          {table.getHeaderGroups().map((headerGroup, key) => (
            <TableRow key={key}>
              {headerGroup.headers.map((column, key2) => (
                <TableHead key={key2} colSpan={column.colSpan}>
                  {flexRender(
                    column.column.columnDef.header,
                    column.getContext(),
                  )}
                </TableHead>
              ))}
            </TableRow>
          ))}
        </TableHeader>
        <TableBody>
          {table.getRowModel().rows.map((row, key) => (
            <TableRow key={key}>
              {row.getVisibleCells().map((cell, key2) => (
                <TableCell key={key2}>
                  {flexRender(cell.column.columnDef.cell, cell.getContext())}
                </TableCell>
              ))}
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <Pagination>
        <PaginationContent>
          <PaginationItem>
            <PaginationPrevious
              onClick={() => table.previousPage()}
              disabled={!table.getCanPreviousPage()}
            />
          </PaginationItem>
          <PaginationItem>
            <PaginationLink>
              {pagination.pageIndex + 1}{" "}
              <span className="text-muted-foreground">/ {pageCount}</span>
            </PaginationLink>
          </PaginationItem>
          <PaginationItem>
            <PaginationNext
              onClick={() => table.nextPage()}
              disabled={!table.getCanNextPage()}
            />
          </PaginationItem>
        </PaginationContent>
      </Pagination>
    </>
  );
}
