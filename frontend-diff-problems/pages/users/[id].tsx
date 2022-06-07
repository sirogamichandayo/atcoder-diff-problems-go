import { useRouter } from "next/router";
import {
  Avatar,
  Box,
  Card,
  CardContent,
  CardMedia,
  CircularProgress,
  Grid,
  LinearProgress,
  Paper,
  Stack,
  Typography,
} from "@mui/material";
import { PieChart } from "recharts";
import { Pie, ResponsiveContainer } from "recharts";

export default function Post() {
  const router = useRouter();
  const { id } = router.query;
  return (
    <>
      <Grid container spacing={2} sx={{ height: 186 }}>
        <Grid item xs={4}>
          <Profile />
        </Grid>
        <Grid item xs={8}>
          <ProblemProgress />
        </Grid>
      </Grid>
    </>
  );
}

function ProblemProgress() {
  return (
    <Card sx={{ height: 186, display: "flex" }}>
      <Grid container alignItems={"center"} justify={"center"}>
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
