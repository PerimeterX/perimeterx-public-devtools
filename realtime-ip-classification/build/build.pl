#!/usr/bin/env perl
use strict;
use warnings;

use Text::CSV_XS;
use Net::Works::Network;
use MaxMind::DB::Writer::Tree;

my %types = (
    service => 'utf8_string',
    region  => 'utf8_string',
);

my $tree = MaxMind::DB::Writer::Tree->new(
    database_type => 'Feed-IP-Data',
    description => { en => 'Amazon IP data' },
    ip_version => 6,
    map_key_type_callback => sub { $types{ $_[0] } },
    record_size => 24,
);

my $file = $ARGV[0] or die "Need to get CSV file on the command line\n";
print "==> ", $file, "\n";
open(my $fh, "<", $file) or die "$file: $!";

my $csv = Text::CSV_XS->new();
$csv->column_names($csv->getline($fh));
while (my $row = $csv->getline($fh)) {
    my $network = Net::Works::Network->new_from_string( string => $row->[2] );
    my $metadata = { service => $row->[0], region => $row->[1] };
    $tree->insert_network($network, $metadata);
}
close $fh;

my $filename = $ARGV[1] or die "Need to get mmdb file on the command line\n";
open my $ofh, '>:raw', $filename;
$tree->write_tree( $ofh );
close $ofh;

print "$filename created\n";