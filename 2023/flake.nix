{ pkgs ? import <nixpkgs> {} }:
let
  my-python-packages = ps: with ps; [
    pytest
    pytest-cache
    pytest-subtests
  ];
  my-python = pkgs.python3.withPackages my-python-packages;
in my-python.env
