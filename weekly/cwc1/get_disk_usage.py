# Enter your code here. Read input from STDIN. Print output to STDOUT
import os
import sys


def list_files_and_sub_files(path):
    """
    prints the files present in the directory and sub-directories of the given path.
    :param path: Path of a directory
    :return:
    """
    files = []
    for (dirpath, dirnames, filenames) in os.walk(path):
        for f in filenames:
            path = os.path.join(dirpath, f)
            file_stats = os.stat(path)
            files.append([path, file_stats.st_size])
    output = {
        "files": files
    }
    print(output)


def list_files(path):
    """
    prints the files present in the directory of the given path.
    :param path: Path of a directory
    :return:
    """
    files = [[f.path, os.stat(f.path).st_size] for f in
             os.scandir(path) if
             f.is_file()]
    output = {
        "files": files
    }
    print(output)


if __name__ == '__main__':
    if not sys.stdin.isatty():
        for line in sys.stdin:
            list_files_and_sub_files(line.rstrip())
