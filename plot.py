import matplotlib.pyplot as plt
from argparse import ArgumentParser
import json


def main():
    parser = ArgumentParser()
    parser.add_argument("filename", type=str)
    parser.add_argument("-e", "--every", type=int, default=1)
    args = parser.parse_args()
    
    with open(args.filename, "r") as file:
        data = json.load(file)
        x, y = data["x"], data["y"]

    plt.figure(figsize=(6,4), dpi=140)
    plt.scatter(x[0], y[0], color="blue", marker="o", s=80, zorder=1, label="start")
    plt.plot(x[::args.every], y[::args.every], color="red", lw=1.5, zorder=0, label=f"{len(x) - 1} steps")
    plt.scatter(x[-1], y[-1], color="green", marker="o", s=80, zorder=1, label="end")
    plt.legend()
    plt.axis("equal")
    plt.show()

if __name__ == "__main__":
    main()
