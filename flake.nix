{
  description = "A simple Go package";

  nixConfig.bash-prompt = "\[nix-develop\]$ ";

  inputs.nixpkgs.url = "nixpkgs/nixpkgs-unstable";

  outputs = { self, nixpkgs }:
    let
      lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";
      version = builtins.substring 0 8 lastModifiedDate;
      supportedSystems = [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });
    in
    {
      devShells = forAllSystems (system:
        let pkgs = nixpkgsFor.${system};
        in {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [
              go
              gotools
              gopls
              go-outline
              gocode
              godef
              golint
              git
              delve
              golangci-lint
            ];
          };
        });

      devShell = forAllSystems (system: self.devShells.${system}.default);

      packages =
        forAllSystems
          (system:
            let
              pkgs = nixpkgsFor.${system};
            in
            {
              alizer-tester = pkgs.buildGoModule {
                pname = "alizer-tester";
                inherit version;
                src = ./.;
                vendorSha256 = "MVSleIXfP6ZyVMl/GGvm2YgT2/oaEkGMSabEw3FX9yU=";
              };
            });
      defaultPackage = forAllSystems (system: self.packages.${system}.alizer-tester);
    };
}