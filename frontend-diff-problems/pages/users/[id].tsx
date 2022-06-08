import { useRouter } from "next/router";
import {
  Box,
  Card,
  CardContent,
  CardMedia,
  CircularProgress,
  Grid,
  LinearProgress,
  Stack,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";
import NextLink from "next/link";
import { Link as MUILink } from "@mui/material";
import React from "react";
import {
  LineChart,
  CartesianGrid,
  Legend,
  Line,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from "recharts";

export default function Post() {
  const router = useRouter();
  const { id } = router.query;
  return (
    <>
      <Grid container spacing={2} sx={{ height: 186 }}>
        <Grid item xs={4}>
          <Profile />
        </Grid>
        <Grid item xs={8} sx={{ minWidth: 460 }}>
          <ProblemProgress />
        </Grid>
        <Grid item xs={4}>
          <NearRankingTable />
        </Grid>
        <Grid item xs={8}>
          <ProblemGraph />
        </Grid>
      </Grid>
    </>
  );
}

function Profile() {
  return (
    <Card sx={{ minWidth: 280, height: 186, display: "flex" }}>
      <CardMedia
        component="img"
        sx={{ width: 120, height: 120 }}
        image={"https://assets.leetcode.com/users/avatars/avatar_1647097868.png"}
        alt="Live from space album cover"
      />
      <Box sx={{ display: "flex", flexDirection: "column", height: 120, width: 140 }}>
        <CardContent sx={{ flex: "1 0 auto" }}>
          <Typography variant="h6" component="div" sx={{ fontWeight: "bold" }}>
            tourist
          </Typography>
          <Typography variant="subtitle1" color="text.secondary" component="div">
            Rating{" "}
            <Box component={"span"} color="text.primary" fontWeight={"fontWeightBold"}>
              4018
            </Box>
          </Typography>
          <Typography variant="subtitle1" color="text.secondary" component="div">
            Rank{" "}
            <Box component={"span"} color="text.primary" fontWeight={"fontWeightBold"}>
              1
            </Box>
            th
          </Typography>
        </CardContent>
      </Box>
    </Card>
  );
}

function ProblemProgress() {
  return (
    <Card sx={{ height: 186, display: "flex" }}>
      <Grid container sx={{ alignItems: "center", justify: "center" }}>
        <Grid item xs={12} sx={{ height: 16 }}>
          Solved Problems
        </Grid>
        <Grid item xs={3} sx={{ textAlign: "center" }}>
          <Box sx={{ position: "relative", display: "inline-flex" }}>
            <CircularProgress variant="determinate" value={90.5} size={100} thickness={3.0} />
            <Box
              sx={{
                top: 0,
                left: 0,
                bottom: 0,
                right: 0,
                position: "absolute",
                display: "flex",
                alignItems: "center",
                justifyContent: "center",
              }}
            >
              <Typography
                variant="caption"
                component="div"
                color="text.secondary"
              >{`90.5%`}</Typography>
            </Box>
          </Box>
        </Grid>
        <Grid item xs={4.5} sx={{ height: 160 }}>
          <Stack sx={{ width: "95%", color: "grey.500" }} spacing={3}>
            <EachProblemProgress />
            <EachProblemProgress />
            <EachProblemProgress />
          </Stack>
        </Grid>
        <Grid item xs={4.5} sx={{ height: 160 }}>
          <Stack sx={{ width: "95%", color: "grey.500" }} spacing={3}>
            <EachProblemProgress />
            <EachProblemProgress />
            <EachProblemProgress />
          </Stack>
        </Grid>
      </Grid>
    </Card>
  );
}

function EachProblemProgress() {
  return (
    <>
      <Stack>
        orange 20/50 40%
        <LinearProgress color="secondary" />
      </Stack>
    </>
  );
}

function NearRankingTable() {
  return (
    <Card>
      <TableContainer>
        <Table aria-label="simple table">
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
      <NextLink href="/ranking" passHref>
        <MUILink variant="body2">All Ranking</MUILink>
      </NextLink>
    </Card>
  );
}

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

const series = [
  {
    name: "diff",
    color: "red",
    data: [
      { date: new Date(2020, 1, 1).getTime(), value: 100 },
      // { date: new Date(2020, 1, 2).getTime(), value: 200 },
      { date: new Date(2020, 1, 3).getTime(), value: 200 },
      { date: new Date(2020, 1, 4).getTime(), value: 300 },
    ],
  },
  {
    name: "rating",
    color: "blue",
    data: [
      { date: new Date(2020, 1, 1).getTime(), value: 200 },
      // { date: new Date(2020, 1, 2).getTime(), value: 200 },
      { date: new Date(2020, 1, 3).getTime(), value: 300 },
      { date: new Date(2020, 1, 4).getTime(), value: 400 },
    ],
  },
];

function ProblemGraph() {
  return (
    <Card sx={{ minWidth: 280, height: 400, display: "flex" }}>
      <ResponsiveContainer width="100%" height="100%">
        <LineChart>
          <CartesianGrid strokeDasharray="3 3" />
          <XAxis
            scale="time"
            domain={["dataMin", "dataMax"]}
            dataKey="date"
            tickFormatter={(unixTime) => new Date(unixTime).toLocaleDateString()}
            type="number"
            allowDuplicatedCategory={false}
          />
          <YAxis dataKey="value" />
          <Legend />
          <Tooltip labelFormatter={(value) => new Date(value).toLocaleDateString()} />
          {series.map((s) => (
            <Line dataKey="value" data={s.data} name={s.name} key={s.name} stroke={s.color} />
          ))}
        </LineChart>
      </ResponsiveContainer>
    </Card>
  );
}
