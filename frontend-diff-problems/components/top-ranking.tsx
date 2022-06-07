import {
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";
import React from "react";

function createData(name: string, rank: number, diffSum: number) {
  return { name, rank, diffSum };
}

const rows = [
  createData("sirogami", 1, 300.0),
  createData("kurogami", 2, 200.0),
  createData("akagami", 3, 100.0),
  createData("chagami", 4, 90.0),
  createData("aaaaaaa", 5, 90.0),
];

export default function TopRanking() {
  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Rank</TableCell>
            <TableCell align="right">UserId</TableCell>
            <TableCell align="right">Difficulty Sum</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {rows.map((row) => (
            <TableRow key={row.name} sx={{ "&:last-child td, &:last-child th": { border: 0 } }}>
              <TableCell>{row.rank}</TableCell>
              <TableCell align={"right"}>{row.name}</TableCell>
              <TableCell align={"right"}>{row.diffSum}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
