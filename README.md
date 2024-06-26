<a name="readme-top"></a>

<h3 align="start">Developing a CPU Scheduler using Go Language</h3>

  <p align="start">
    This project developed a user-friendly terminal program written in Go to analyze and visualize CPU scheduling algorithms. Users can select from five popular algorithms (FCFS, SJF, SRTF, Priority, RR) and the program calculates performance metrics (average waiting time, turnaround time) for each.

[![Product Name Screen Shot][product-screenshot]](https://example.com)

Key features include:

1. User selection of scheduling algorithms
2. CSV data import for streamlined analysis
3. Process table and Gantt chart visualizations for each algorithm

Go was chosen for its superior speed compared to Python and Java, resulting in faster analysis and visualizations. The project addressed challenges in balancing Go's speed with visualization needs by opting for hand-coded visualization logic to maintain performance.
<br />

<a href="https://github.com/supermonmon/cpu-scheduling-visualizer-go"><strong>Explore the docs »</strong></a>
<br />
<br />
<a href="https://github.com/supermonmon/cpu-scheduling-visualizer-go/issues">Request Feature</a>
·
<a href="https://github.com/supermonmon/cpu-scheduling-visualizer-go/issues/new">Report Bug</a>

  </p>
</div>

<!-- GETTING STARTED -->

## Getting Started

### Installation

Go 1.22.1 +

```sh
go get https://github.com/supermonmon/cpu-scheduling-visualizer-go
```

### Usage

```sh
go run main.go ./data/data.csv
```

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<!-- MARKDOWN LINKS & IMAGES -->

[product-screenshot]: images/snapshot.jpg
