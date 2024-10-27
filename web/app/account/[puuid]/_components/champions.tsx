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
import { ChampionMastery } from "@/masteries/v1/masteries_pb";
import { PlainMessage } from "@bufbuild/protobuf";

type Props = {
  masteries: PlainMessage<ChampionMastery>[];
};

const columnHelper = createColumnHelper<ChampionMastery>();
const defaultColumns = [
  columnHelper.accessor("championId", {
    header: "Champion",
    cell: (props) => props.row.original.championId,
  }),
  columnHelper.accessor("championLevel", {
    header: "Gespielt",
  }),
  columnHelper.accessor("championLevel", {
    header: "KDA",
  }),
  columnHelper.accessor("championLevel", {
    header: "Gold",
  }),
  columnHelper.accessor("championLevel", {
    header: "CS",
  }),
  columnHelper.accessor("championLevel", {
    header: "Durchschnittlich verursachter Schaden",
  }),
  columnHelper.accessor("championLevel", {
    header: "Durchschnittlich erhaltener Schaden",
  }),
  columnHelper.accessor("championLevel", {
    header: "Zweifachtötung",
  }),
  columnHelper.accessor("championLevel", {
    header: "Dreifachtötung",
  }),
  columnHelper.accessor("championLevel", {
    header: "Vierfachtötung",
  }),
  columnHelper.accessor("championLevel", {
    header: "Fünffachtötung",
  }),
];

export function Champions(props: Props) {
  const [pagination, setPagination] = useState<PaginationState>({
    pageIndex: 0,
    pageSize: 10,
  });

  const pageCount = useMemo(
    () => Math.ceil(props.masteries.length / pagination.pageSize),
    [pagination.pageSize, props.masteries.length],
  );

  const table = useReactTable({
    columns: defaultColumns,
    data: props.masteries.sort((a, b) => b.championLevel - a.championLevel),
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
