# Day 07

> For a more comprehensive article on this day, please refer to this [link.](https://github.com/MichaelCade/90DaysOfDevOps/blob/main/2022/Days/day07.md)

I have compressed a few days into this one day. Although, I did not do everything written here on this day alone. It 
spanned over a few days. I have also added a few things I learnt on my own.

## Why do we need to learn a programming language as DevOps Engineers?

We need to learning a programming language because some of the tools we will use for most of the DevOps practiceswe will embark on are written in various programming languages. We do not now need to learn all of the wide array of functionalities of each programming language we will need to use, we just need to understand or have an indepth foundational knowledge of the language so we can understand other codes, and also when we need it to perform certain repetitive tasks, we can easily employ the programming languages to help perform those tasks on our behalf.

## Which Programming Language Will We Use?

Most DevOps engineers make use of Python as their preferred programming language simply because it is an easy to understand and use interpreted language that can be used to perform various repetitive tasks, and it has a very vast library of various modules developed by a large community of maintainers. For every task you may need to perform, it is very likely that Python has a corresponding module to help with that task. Following [@MichaelCade's 90DaysOfDevOps](https://github.com/MichaelCade/90DaysOfDevOps/blob/main/2022/Days/day07.md), we will be starting with Golang which is a statically tpyed compiled language used to build CloudNative tools used by DevOps Engineers e.g. Terraform, Prometheus, Docker, K8s, Grafana, etc.

Another Benefit of Golang is that unlike Python which requires external libraries to be installed when runing on a remote machine, it is compiled into a single binary which can be run anywhere. Golang is also platform independent.

> Installing [Golang](https://dev.to/deadwin19/how-to-install-golang-on-wslwsl2-2880) on wsl2

``` SHELL
# Click the hyperlink above to get the https link of the latest version.
wget https://go.dev/dl/go1.21.0.linux-amd64.tar.gz

# Unzip the archive and store it in a file
sudo tar -xvf go1.21.0.linux-amd64.tar.gz

# Move the go file to the /usr/local directory
sudo mv go /usr/local

# Add the file to the PATH environment variable
export PATH=$PATH:/usr/local/go/bin

# Check the installed go version
go version
```

### Resources

- [Jake Wright - Learn Go in 12 minutes](https://www.youtube.com/watch?v=C8LgvuEBraI)
- [TechworldWithNana - Golang full course - 3 hours 24 mins](https://www.youtube.com/watch?v=yyUHQIec83I)
